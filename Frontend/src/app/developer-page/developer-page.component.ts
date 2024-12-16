import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserModel } from 'src/models/user-model';
import { UserServiceService } from 'src/services/user-service.service';

@Component({
  selector: 'developer-page',
  templateUrl: './developer-page.component.html',
  styleUrls: ['./developer-page.component.css', '../../styles.css'],
})
export class DeveloperPageComponent {
  public isDeveloper = false;
  public userList : UserModel[] = [];

  constructor(private route: ActivatedRoute,private userService: UserServiceService) {}

  ngOnInit(): void {
    this.route.params.subscribe((params) => {
      this.isDeveloper = params['isDevelopmentEnvironment'];
      this.userService.getAllUsers().subscribe(result => {
        this.userList = result;
      });
    });
  }
}
