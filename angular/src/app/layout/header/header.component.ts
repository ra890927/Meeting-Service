import { Component, inject } from '@angular/core';

// angular material import
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { Router, RouterLink, RouterLinkActive } from '@angular/router';
import { UserService } from '../../API/user.service';
import { AuthService } from '../../API/auth.service';
@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    MatToolbarModule,
    MatIconModule,
    MatButtonModule, 
    RouterLink, 
    RouterLinkActive
  ],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent {
  userservice = inject(UserService);
  authservice = inject(AuthService);
  router = inject(Router);
  constructor() { }
  logout() {
    console.log('logout');
    this.authservice.logout().subscribe(
      response => {
        console.log(response);
        if (response.status === 'success') {
          console.log('logout success');
          
          this.userservice.clean();
          this.router.navigate(['']);
        }
      },
      error => {
        console.log(error);
      }
    );
  }
}
