import { Component } from '@angular/core';

import { UserComponent } from './user/user.component';
import { HeaderComponent } from '../layout/header/header.component';
import { RoomComponent } from './room/room.component';
import { MatTabsModule } from '@angular/material/tabs';

@Component({
  selector: 'app-monitor',
  standalone: true,
  imports: [
    HeaderComponent,
    UserComponent,
    RoomComponent,
    MatTabsModule,
  ],
  templateUrl: './monitor.component.html',
  styleUrl: './monitor.component.css'
})
export class MonitorComponent {

}
