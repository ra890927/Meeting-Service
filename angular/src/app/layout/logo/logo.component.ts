import { Component,inject } from '@angular/core';
import { Router } from '@angular/router';
@Component({
  selector: 'app-logo',
  standalone: true,
  imports: [],
  templateUrl: './logo.component.html',
  styleUrl: './logo.component.css'
})
export class LogoComponent {
  constructor(private route: Router) {
  }
  navigate() {
    this.route.navigate(['/homepage']);
  }
}
