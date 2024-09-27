import { Component } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';
import { UserModel } from 'src/models/user-model';
import { UserServiceService } from 'src/services/user-service.service';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent {
  public loginForm: FormGroup = new FormGroup({});
  public selectedUser = {
    id: 0,
    firstName: "Please choose a name...",
    lastName: "",
    username: ""
  } as UserModel;
  public availableUsers: UserModel[] = [];
  public currentSearchUser: number = 0;

  constructor (private toaster: ToastrService, private userService: UserServiceService) {}

  ngOnInit(): void {
    this.userService.getAllUsers().subscribe(result => {
      this.availableUsers = result;
    });
  }

  public changeSelectedUser(user: any): void {
    const newUser = this.availableUsers.find(x => x.username === user.target.value) as UserModel;
    this.selectedUser = newUser;
  }

  // this is searching by ID, not by e-mail
  public searchForUser(): void {
    this.userService.getUser(this.currentSearchUser).subscribe(result => {
      console.log('results: ', result);
      this.selectedUser = result;
    });
  }


}
