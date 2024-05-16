import { Component } from '@angular/core';
import { HeaderComponent } from '../layout/header/header.component';
import { RoomSchedulerComponent } from '../room-scheduler/room-scheduler.component';

@Component({
  selector: 'app-reservation',
  standalone: true,
  imports: [
    HeaderComponent,
    RoomSchedulerComponent
  ],
  templateUrl: './reservation.component.html',
  styleUrl: './reservation.component.css'
})
export class ReservationComponent {

}
