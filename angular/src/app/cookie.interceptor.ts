import { HttpInterceptorFn, HttpXsrfTokenExtractor } from '@angular/common/http';
import {inject} from '@angular/core';
export const cookieInterceptor: HttpInterceptorFn = (req, next) => {
  req = req.clone({
    // approve the cookie from the cross-origin request
    withCredentials: true,
    setHeaders: {
      'sameSite': 'None',
      'Authorization': String(window.sessionStorage.getItem('token'))
    }

  });
  console.log('cookieInterceptor');
  return next(req);
};
