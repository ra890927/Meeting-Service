import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { Router } from '@angular/router';
import { FooterComponent } from '../layout/footer/footer.component';
import { LogoComponent } from '../layout/logo/logo.component';


@Component({
  selector: 'app-homepage',
  standalone: true,
  imports: [
    MatButtonModule,
    FooterComponent,
    LogoComponent
  ],
  templateUrl: './homepage.component.html',
  styleUrl: './homepage.component.css'
})
export class HomepageComponent {
  constructor(private router: Router) { }

  ngOnInit(): void {
  }
  navigate(path: string) {
    this.router.navigate([path]); // navigate to the path
  }

}
