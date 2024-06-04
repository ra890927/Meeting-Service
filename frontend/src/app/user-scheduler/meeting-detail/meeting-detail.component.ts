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
import { FileService } from '../../API/file.service';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
interface Tag{
  id: number;
  tag: string;
  description: string;
  codeTypeId: number;
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
interface fileFormat{
  file_id: string;
  file_name: string;
  url?: string;
  Uploader_id?: number;
}
@Component({
  selector: 'app-meeting-detail',
  standalone: true,
  imports: [MatButtonModule,DatePipe,MatTooltipModule, MatDialogClose, MatChipsModule, MatIconModule, MatMenuModule],
  templateUrl: './meeting-detail.component.html',
  styleUrl: './meeting-detail.component.css'
})
export class MeetingDetailComponent {
  uploadedFiles: fileFormat[] = [];
  constructor(public dialogRef: MatDialogRef<MeetingDetailComponent>,
    @Inject(MAT_DIALOG_DATA) public data: MeetingDetail ,private file: FileService) {
      console.log("dialog meeting:",data);
      if(this.data.id !== undefined){
        this.file.getFileByMeetingId(this.data.id).subscribe((response: any) => {
          if(response.status === 'success'){
            if(response.data.length  !== null){
              this.uploadedFiles = response.data;
              console.log(this.uploadedFiles);
            }
            else
              console.log('No file found');
          }
          else{
            console.log('Get file failed');
          }
        }
      );
    }
    }
    download(file: fileFormat){
      if(file.url){
        const downloadUrl = `${file.url}`;
        window.open(downloadUrl);
      }
    }

}
