import { Component, Inject } from "@angular/core";
import { MatButtonModule } from "@angular/material/button";
import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from "@angular/material/dialog";

@Component({
    selector: './delete-alarm',
    templateUrl: './delete-alarm.html',
    styleUrl: './delete-alarm.css',
    standalone: true,
    imports: [
      // MatButtonModule, MatDialogActions, MatDialogClose, MatDialogTitle, MatDialogContent,
      MatDialogModule,
      MatButtonModule,
    ],
  })
  export class DeleteAlarm {
    constructor(public dialogRef: MatDialogRef<DeleteAlarm>, @Inject(MAT_DIALOG_DATA) public data: any) {}
    class = this.data.class;
  
    delete(): void {
      this.dialogRef.close({isDelete: true, deleteClassType: this.data.deleteClassType});
    }
  }