import { Component, inject, signal } from '@angular/core';

// angular material import
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { ActivatedRoute, Router, RouterLink, RouterLinkActive } from '@angular/router';
import { UserService } from '../../API/user.service';
import { AuthService } from '../../API/auth.service';
import { MatTooltipModule } from '@angular/material/tooltip';
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    MatToolbarModule,
    MatIconModule,
    MatButtonModule, 
    RouterLink, 
    RouterLinkActive,
    MatTooltipModule,
    CommonModule
  ],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent {
  userservice = inject(UserService);
  authservice = inject(AuthService);
  router = inject(Router);
  isAdmin = false;
  currentUrl: string | undefined;
  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.currentUrl = this.route.snapshot.url.join('/');
    this.authservice.whoami().subscribe(
      response => {
        console.log(response);
        if (response.data.user.role === 'admin') {
          this.isAdmin = true;
        }
      },
      error => {
        console.log(error);
      }
    );
  }

  // navigate to the admin page
  navigate(path: string) {
    this.router.navigate([path]);
    
  }

  adminSwitch() {
    if (this.currentUrl === 'monitor') {
      this.navigate('/profile');
    }
    else {
      this.navigate('/monitor');
    }
  }


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
