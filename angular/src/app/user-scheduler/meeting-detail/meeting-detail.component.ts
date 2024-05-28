import {Component, Inject} from '@angular/core';
import { DatePipe } from '@angular/common';
import {
  MatDialog,
  MAT_DIALOG_DATA,
  MatDialogRef,
  MatDialogTitle,
  MatDialogContent,
  MatDialogActions,
  MatDialogClose,
} from '@angular/material/dialog';
import {MatButtonModule} from '@angular/material/button';
import {MatChipsModule} from '@angular/material/chips';
import {MatTooltipModule} from '@angular/material/tooltip';
interface Tag{
  type_name: string;
  type_desc: string;
}

interface RoonInfo{
  room_name: string;
  rules: Tag[];
}
interface User{
  id: number;
  username: string;
  email: string;
}
interface MeetingDetail{
  id: string;
  title: string;
  description: string;
  participants: User[];
  start: string;
  end: string;
  Organizer: User;
  RoomDetail: RoonInfo;
}

@Component({
  selector: 'app-meeting-detail',
  standalone: true,
  imports: [MatButtonModule,DatePipe,MatTooltipModule, MatDialogClose, MatChipsModule],
  templateUrl: './meeting-detail.component.html',
  styleUrl: './meeting-detail.component.css'
})
export class MeetingDetailComponent {
  constructor(public dialogRef: MatDialogRef<MeetingDetailComponent>,
    @Inject(MAT_DIALOG_DATA) public data: MeetingDetail) {
      console.log(data);
    }

}
