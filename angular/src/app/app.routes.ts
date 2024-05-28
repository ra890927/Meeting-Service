import { Routes } from '@angular/router';
import { HomepageComponent } from './homepage/homepage.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { ReservationComponent } from './reservation/reservation.component';
import { RoomSchedulerComponent } from './room-scheduler/room-scheduler.component';
import { MonitorComponent } from './monitor/monitor.component';
import { UserSchedulerComponent } from './user-scheduler/user-scheduler.component';

export const routes: Routes = [
    { path:  "", component: HomepageComponent},
    { path: "login", component: LoginComponent},
    { path: "register", component: RegisterComponent},
    { path: "reservation", component: ReservationComponent},
    { path: "room-scheduler", component: RoomSchedulerComponent},
    { path: "monitor", component: MonitorComponent},
    { path: "user-scheduler", component: UserSchedulerComponent},
];
