import { Component, Inject} from '@angular/core';
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

@Component({
  selector: 'app-pop-up-delete-confirm',
  standalone: true,
  imports: [MatButtonModule,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions,
    MatDialogClose,],
  templateUrl: './pop-up-delete-confirm.component.html',
  styleUrl: './pop-up-delete-confirm.component.css'
})
export class PopUpDeleteConfirmComponent {
  
  constructor(
    public dialogRef: MatDialogRef<PopUpDeleteConfirmComponent>
  ) {}
  onNoClick(): void {
    this.dialogRef.close();
  }
}
