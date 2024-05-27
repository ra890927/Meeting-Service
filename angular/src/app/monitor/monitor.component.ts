import { Component } from '@angular/core';

// import components
import { HeaderComponent } from '../layout/header/header.component';
import { UserComponent } from './user/user.component';
import { RoomComponent } from './room/room.component';
import { TagComponent } from './tag/tag.component';

// import modules
import { MatTabsModule } from '@angular/material/tabs';


@Component({
  selector: 'app-monitor',
  standalone: true,
  imports: [
    HeaderComponent,
    UserComponent,
    RoomComponent,
    TagComponent,
    MatTabsModule,
  ],
  templateUrl: './monitor.component.html',
  styleUrl: './monitor.component.css'
})
export class MonitorComponent {

}
