import { Component, ElementRef, ViewChild } from '@angular/core';
import { users } from './users';

import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';

import { HeaderComponent } from '../layout/header/header.component';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { MatInput } from '@angular/material/input';
import { UserComponent } from './user/user.component';

@Component({
  selector: 'app-monitor',
  standalone: true,
  imports: [
    HeaderComponent,
    UserComponent,
    MatTabsModule,
    MatCardModule,
    MatListModule,
    MatIconModule,
    CommonModule

  ],
  templateUrl: './monitor.component.html',
  styleUrl: './monitor.component.css'
})
export class MonitorComponent {

}
