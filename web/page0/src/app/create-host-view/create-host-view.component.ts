import {Component} from '@angular/core';
import {FormBuilder} from '@angular/forms';
import {ActivatedRoute, NavigationEnd, Router} from '@angular/router';
import {HostService} from '../host.service';
import {MatSnackBar} from '@angular/material/snack-bar';
import {BehaviorSubject, Subject} from 'rxjs';
import {
  filter,
  map,
  mergeMap,
  shareReplay,
  switchMap,
  take,
  takeUntil,
  tap,
  withLatestFrom,
} from 'rxjs/operators';

import {Store} from 'src/app/store/store';
import {
  runtimeSelectorFactory,
  runtimesLoadStatusSelector,
} from '../store/selectors';
import {RuntimeViewStatus} from '../interface/runtime-interface';
import {DEFAULT_HOST_SETTING, DEFAULT_ZONE} from '../settings';
import {ResultType} from '../interface/result-interface';

@Component({
  standalone: false,
  selector: 'app-create-host-view',
  templateUrl: './create-host-view.component.html',
  styleUrls: ['./create-host-view.component.scss'],
})
export class CreateHostViewComponent {
  constructor(
    private hostService: HostService,
    private formBuilder: FormBuilder,
    private router: Router,
    private snackBar: MatSnackBar,
    private activatedRoute: ActivatedRoute,
    private store: Store
  ) {
    this.queryParams$ = this.router.events.pipe(
      filter((event): event is NavigationEnd => event instanceof NavigationEnd),
        mergeMap(() => this.activatedRoute.queryParams),
        shareReplay(1)
    );
    this.runtime$ = this.queryParams$.pipe(
      map(params => (params['runtime'] as string) ?? ''),
        switchMap(alias =>
                  this.store.select(runtimesLoadStatusSelector).pipe(
                    filter(status => status === RuntimeViewStatus.done),
                      switchMap(() =>
                                this.store.select(runtimeSelectorFactory({alias})).pipe(
                                  map(runtime => {
                                    if (!runtime) {
                                      throw new Error(`No runtime of alias ${alias}`);
                                    }
                                    return runtime;
                                  })
                                )
                               )
                  )
                 )
    );
    this.previousUrl$ = this.queryParams$.pipe(
      map(params => (params['previousUrl'] as string) ?? 'list-runtime')
    );
    this.zones$ = this.runtime$.pipe(
      map(runtime => runtime.zones),
        tap(zones => {
        if (zones?.includes(DEFAULT_ZONE)) {
          this.hostForm!.controls.zone.setValue(DEFAULT_ZONE);
        }
      })
    );
    this.status$ = new BehaviorSubject<string>('done');
    this.hostForm = this.formBuilder.group({
      zone: ['ap-northeast2-a'],
      machine_type: [DEFAULT_HOST_SETTING.gcp?.machine_type],
      min_cpu_platform: [DEFAULT_HOST_SETTING.gcp?.min_cpu_platform],
    });
    this.queryParams$.pipe(takeUntil(this.ngUnsubscribe)).subscribe();
  }

  private ngUnsubscribe = new Subject<void>();

  queryParams$;
  runtime$;
  previousUrl$;
  zones$;

  ngOnDestroy() {
    this.ngUnsubscribe.next();
    this.ngUnsubscribe.complete();
  }

  status$;

  hostForm;

  // TODO: refactor with 'host status'
  showProgressBar(status: string | null) {
    return status === 'sending create request';
  }

  onSubmit() {
    const {machine_type, min_cpu_platform, zone} = this.hostForm.value;
    if (!machine_type || !min_cpu_platform || !zone) {
      return;
    }

    this.runtime$
      .pipe(
        take(1),
        withLatestFrom(this.previousUrl$),
        takeUntil(this.ngUnsubscribe),
        switchMap(([runtime, previousUrl]) => {
          this.status$.next('sending create request');
          return this.hostService
            .createHost(
              {
                gcp: {
                  machine_type,
                  min_cpu_platform,
                },
              },
              runtime,
              zone
            )
            .pipe(
              map(result => ({
                result,
                previousUrl,
              }))
            );
        })
      )
      .subscribe({
        next: ({result, previousUrl}) => {
          if (result.type === ResultType.waitStarted) {
            this.status$.next('done');
            this.router.navigate([previousUrl]);
            this.snackBar.dismiss();
          }
        },
        error: error => {
          this.snackBar.open(error.message, 'Dismiss');
        },
      });
  }

  onCancel() {
    this.previousUrl$
      .pipe(takeUntil(this.ngUnsubscribe))
      .subscribe(previousUrl => {
        this.router.navigate([previousUrl]);
      });
  }
}
