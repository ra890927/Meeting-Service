import { Component, Inject} from '@angular/core';
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
import {FormsModule} from '@angular/forms';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
export interface DialogData {
  title: string;
  description: string;
  startTime: string;
  endTime: string;
}
@Component({
  selector: 'app-pop-up-details',
  standalone: true,
  imports: [MatFormFieldModule,
    MatInputModule,
    FormsModule,
    MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatDialogClose,DatePipe, MatIconModule],
  templateUrl: './pop-up-details.component.html',
  styleUrl: './pop-up-details.component.css'
})
export class PopUpDetailsComponent {
  constructor(
    public dialogRef: MatDialogRef<PopUpDetailsComponent>,
    @Inject(MAT_DIALOG_DATA) public data: DialogData,
  ) {}
  onNoClick(): void {
    this.dialogRef.close();
  }
}
