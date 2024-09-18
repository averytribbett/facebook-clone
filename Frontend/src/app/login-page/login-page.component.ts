import { Component } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent {
  public loginForm: FormGroup = new FormGroup({});

  constructor (private toaster: ToastrService) {}
  
  ngOnInit(): void {
    this.initForm();
  }

  public initForm(): void {
    this.loginForm = new FormGroup({
      username: new FormControl(''),
      password: new FormControl('')
    });
  }

  public login(): void {
    const msg = "Current user has attempted to log in with username " + this.loginForm.get('username')?.value + " and password " + this.loginForm.get('password')?.value;
    this.toaster.success(msg);
  }


}
