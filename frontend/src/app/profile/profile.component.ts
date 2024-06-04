import { Component } from '@angular/core';
import { UserSchedulerComponent } from '../user-scheduler/user-scheduler.component';
import { HeaderComponent } from '../layout/header/header.component';
import { FooterComponent } from '../layout/footer/footer.component';

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [
    UserSchedulerComponent,
    HeaderComponent,
    FooterComponent
  ],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css'
})
export class ProfileComponent {
  
}
