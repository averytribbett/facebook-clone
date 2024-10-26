import { Component, EventEmitter, Inject, inject, Output } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-profile-incomplete-warning',
  templateUrl: './profile-incomplete-warning.component.html',
  styleUrls: ['./profile-incomplete-warning.component.css'],
})
export class ProfileIncompleteWarningComponent {
  @Output() returnHomeEmitter$ = new EventEmitter<boolean>();
  readonly dialogRef = inject(MatDialogRef<ProfileIncompleteWarningComponent>);

  constructor(@Inject(MAT_DIALOG_DATA) public data: any) {}

  onNoClick(): void {
    this.dialogRef.close(false);
  }

  cancelProfileCreation(): void {
    this.dialogRef.close(true);
  }
}
