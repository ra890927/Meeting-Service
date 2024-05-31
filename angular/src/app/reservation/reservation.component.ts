import { Component } from '@angular/core';
import { HeaderComponent } from '../layout/header/header.component';
import { RoomSchedulerComponent } from '../room-scheduler/room-scheduler.component';
import { F } from '@angular/cdk/keycodes';
import { FooterComponent } from '../layout/footer/footer.component';
@Component({
  selector: 'app-reservation',
  standalone: true,
  imports: [
    HeaderComponent,
    RoomSchedulerComponent,
    FooterComponent
  ],
  templateUrl: './reservation.component.html',
  styleUrl: './reservation.component.css'
})
export class ReservationComponent {

}
