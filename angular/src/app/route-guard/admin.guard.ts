import { CanActivateFn } from '@angular/router';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../API/auth.service';
import { map } from 'rxjs';
export const adminGuard: CanActivateFn = (route, state) => {
  const Auth = inject(AuthService);
  const router = inject(Router);
  return Auth.whoami().pipe(map((data) => {
    if (data.role == 'admin') {
      return true;
    } else {
      router.navigate(['/login']);
      console.log('You are not an admin!');
      return false;
    }
  }));
};
