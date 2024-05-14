import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { Router } from '@angular/router';

@Component({
  selector: 'app-homepage',
  standalone: true,
  imports: [
    MatButtonModule
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
