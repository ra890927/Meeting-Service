import { CanActivateFn } from '@angular/router';
import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../API/auth.service';
import { map } from 'rxjs/operators';
import { Observable } from 'rxjs';
export const adminGuard: CanActivateFn = (route, state) => {
  const auth = inject(AuthService);
  const router = inject(Router);

  return auth.whoami().pipe(
    map(response => {
      if (response.data.user.role === 'admin') {
        return true;
      } else {
        router.navigate(['/login']);
        return false;
      }
    })
  );
};

