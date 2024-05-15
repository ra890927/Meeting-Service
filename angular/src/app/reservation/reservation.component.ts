import { Component } from '@angular/core';
import { HeaderComponent } from '../layout/header/header.component';
import { ScheduleModule } from '@syncfusion/ej2-angular-schedule';
import { DayService, WeekService, WorkWeekService, MonthService, AgendaService } from '@syncfusion/ej2-angular-schedule';
import { EventSettingsModel} from '@syncfusion/ej2-angular-schedule';

@Component({
  selector: 'app-reservation',
  standalone: true,
  imports: [
    HeaderComponent,
    ScheduleModule
  ],
  providers: [DayService, WeekService, WorkWeekService, MonthService, AgendaService],
  templateUrl: './reservation.component.html',
  styleUrl: './reservation.component.css'
})
export class ReservationComponent {
  public startHour: string = '08:00';
  public endHour: string = '24:00';
  public data: object[] = [{
    Id: 1,
    Subject: 'Meeting',
    StartTime: new Date(2024, 4, 12, 10, 0),
    EndTime: new Date(2024, 4, 12, 12, 30)
  }];
  public eventSettings: EventSettingsModel = {
    dataSource: this.data
  }

}
