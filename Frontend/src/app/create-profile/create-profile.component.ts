import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { UserModel } from 'src/models/user-model';
import { UserServiceService } from 'src/services/user-service.service';
import { MatDialog} from '@angular/material/dialog';
import {MatButton} from '@angular/material/button';
import {MatTooltip} from '@angular/material/tooltip';
import { AuthService } from '@auth0/auth0-angular';
import { ProfileIncompleteWarningComponent } from './profile-incomplete-warning/profile-incomplete-warning.component';

@Component({
  selector: 'app-create-profile',
  templateUrl: './create-profile.component.html',
  styleUrls: ['./create-profile.component.css'],
})
export class CreateProfileComponent {
  @Input() currentUser: String = "";
  @Output() myEmitter$ = new EventEmitter<boolean>();
  public createProfileForm: FormGroup = new FormGroup({});

  constructor(
    private userService: UserServiceService,
    public auth: AuthService,
    public dialogRef: MatDialog
  ){}
  
  ngOnInit(): void {
    this.initForm();
    this.disableUserNameField();
  }

  public initForm(): void {
    this.createProfileForm = new FormGroup({
      firstName: new FormControl(''),
      lastName: new FormControl(''),
      bio: new FormControl(''),
      username: new FormControl(''),
    });
  }

  sendNewUserToBackend(): void {
    const userToSend = {
      firstName: this.createProfileForm.get('firstName')?.value,
      lastName: this.createProfileForm.get('lastName')?.value,
      bio: this.createProfileForm.get('bio')?.value,
      username: this.currentUser ?? this.createProfileForm.get('username')?.value,
    } as UserModel;

    console.log('auth user: ', this.auth.user$);

    this.userService.addNewUser(userToSend).subscribe(result => {
      console.log('the result: ', result);
    });
  }

  public openProfileIncompleteDialog(): void {
    const myDialog = this.dialogRef.open(ProfileIncompleteWarningComponent, {
      width: '250px',
      disableClose: true,
    });
    myDialog.afterClosed().subscribe((result) => {
      this.dialogRef.closeAll();
      if (result) {
        this.confirmReturnHome();
      }
    });
  }

  public confirmReturnHome(): void {
    this.myEmitter$.emit(true);
  }

  public disableUserNameField(): void {
    console.log('current user: ', this.currentUser);
    console.log(this.currentUser !== null && this.currentUser !== "");
    if (this.currentUser !== null && this.currentUser !== "") {
      this.createProfileForm.get('username')?.disable();
    }
  }
}
