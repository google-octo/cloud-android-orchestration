import {Component} from '@angular/core';
import {MatSnackBar} from '@angular/material/snack-bar';
import {Router} from '@angular/router';
import {BehaviorSubject} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import {EnvFormService} from '../env-form.service';
import {EnvService} from '../env.service';
import {ResultType} from '../interface/result-interface';
import {validRuntimeListSelector} from '../store/selectors';
import {Store} from '../store/store';
import {AUTO_CREATE_HOST} from '../utils';
@Component({
  standalone: false,
  selector: 'app-create-env-view',
  templateUrl: './create-env-view.component.html',
  styleUrls: ['./create-env-view.component.scss'],
})
export class CreateEnvViewComponent {
  autoCreateHostToken = AUTO_CREATE_HOST;

  envForm;

  constructor(
    private router: Router,
    private snackBar: MatSnackBar,
    private envService: EnvService,
    private envFormService: EnvFormService,
    private store: Store
  ) {
    this.envForm = this.envFormService.getEnvForm();
    this.runtimes$ = this.store
    .select(validRuntimeListSelector)
    .pipe(map(runtimes => runtimes.map(runtime => runtime.alias)));
    this.zones$ = this.envFormService.getZones$();
    this.hosts$ = this.envFormService.getHosts$();
    this.status$ = new BehaviorSubject<string>('done');
    this.hint$ = this.envForm.controls.host.valueChanges.pipe(
      startWith(this.envForm.controls.host.value),
      map(host => {
        if (host === AUTO_CREATE_HOST) {
          return 'Auto Create may not be completed if you leave Page 0';
        }

        return '';
      })
    );
  }

  runtimes$;
  zones$;
  hosts$;
  status$;
  hint$;

  showProgressBar(status: string | null) {
    return status === 'sending create request';
  }

  onClickAddDevice() {
    this.envFormService.addDevice();
  }

  onClickRegisterRuntime() {
    this.router.navigate(['/register-runtime'], {
      queryParams: {
        previousUrl: 'create-env',
      },
    });
  }

  onClickCreateHost() {
    this.router.navigate(['/create-host'], {
      queryParams: {
        previousUrl: 'create-env',
        runtime: this.envForm.value.runtime,
      },
    });
  }

  onSubmit() {
    const {runtime, zone, host, canonicalConfig} = this.envForm.value;

    this.status$.next('sending create request');
    this.envService.createEnv(runtime, zone, host, canonicalConfig).subscribe({
      next: result => {
        if (result.type === ResultType.waitStarted) {
          this.status$.next('done');
          this.snackBar.dismiss();
          this.router.navigate(['/']);
          this.envFormService.clearForm();
        }
      },
      error: error => {
        this.status$.next('error');
        this.snackBar.open(error.message, 'Dismiss');
      },
    });
  }

  onCancel() {
    this.status$.next('done');
    this.router.navigate(['/']);
    this.envFormService.clearForm();
  }
}
