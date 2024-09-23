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
    name: "Please choose a name...",
    age: 0,
    homeTown: "",
    job: "",
    username: ""
  } as UserModel;
  public availableUsers: UserModel[] = [];

  constructor (private toaster: ToastrService, private userService: UserServiceService) {}

  ngOnInit(): void {
    this.userService.getAllUsers().subscribe(result => {
      this.availableUsers = result;
    });
  }

  public changeSelectedUser(user: any): void {
    const newUser = this.availableUsers.find(x => x.name === user.target.value) as UserModel;
    this.selectedUser = newUser;
  }

}
