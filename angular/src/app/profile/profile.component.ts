import { Component } from '@angular/core';
import { UserSchedulerComponent } from '../user-scheduler/user-scheduler.component';
import { HeaderComponent } from '../layout/header/header.component';

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [
    UserSchedulerComponent,
    HeaderComponent
  ],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css'
})
export class ProfileComponent {
  
}
