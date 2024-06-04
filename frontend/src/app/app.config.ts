import { ApplicationConfig} from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { MatDialogModule } from '@angular/material/dialog';
import { provideHttpClient,withInterceptors} from '@angular/common/http';
import { cookieInterceptor } from './cookie.interceptor';
export const appConfig: ApplicationConfig = {
  providers: [provideRouter(routes), provideAnimationsAsync(), MatDialogModule,provideHttpClient(withInterceptors([cookieInterceptor]))],
};
