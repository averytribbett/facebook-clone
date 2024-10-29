import { Component, EventEmitter, Inject, Input, Output } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { UserModel } from 'src/models/user-model';
import { UserServiceService } from 'src/services/user-service.service';
import { MatDialog } from '@angular/material/dialog';
import { AuthService } from '@auth0/auth0-angular';
import { ProfileIncompleteWarningComponent } from './profile-incomplete-warning/profile-incomplete-warning.component';
import { ActivatedRoute, Router } from '@angular/router';
import { DOCUMENT } from '@angular/common';

@Component({
  selector: 'app-create-profile',
  templateUrl: './create-profile.component.html',
  styleUrls: ['./create-profile.component.css'],
})
export class CreateProfileComponent {
  @Input() currentUser = '';
  @Output() myEmitter$ = new EventEmitter<boolean>();
  @Output() returnHome$ = new EventEmitter<boolean>();
  public createProfileForm: FormGroup = new FormGroup({});
  public isDeveloperMode = false;

  constructor(
    private userService: UserServiceService,
    public auth: AuthService,
    public dialogRef: MatDialog,
    private route: ActivatedRoute,
    private router: Router,
    @Inject(DOCUMENT) public document: Document,
  ) {}

  ngOnInit(): void {
    this.initForm();
    this.route?.params.subscribe((params) => {
      this.isDeveloperMode = params['isDeveloperMode'];
    });
    this.userService.userAuth.user$.subscribe((result) => {
      if (result) {
        this.currentUser = result.name as string;
        this.disableUserNameField();
      }
    });
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
      username:
        this.currentUser ?? this.createProfileForm.get('username')?.value,
    } as UserModel;

    this.userService.addNewUser(userToSend).subscribe(() => {
      this.router.navigate(['/home']);
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
        this.auth.logout({
          logoutParams: {
            returnTo: this.document.location.origin,
          },
        });
      }
    });
  }

  public disableUserNameField(): void {
    if (this.currentUser !== null && this.currentUser !== '') {
      this.createProfileForm.get('username')?.disable();
    }
  }
}
