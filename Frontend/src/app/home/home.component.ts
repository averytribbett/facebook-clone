import { DOCUMENT } from '@angular/common';
import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { UserServiceService } from 'src/services/user-service.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css','../../styles.css'],
})
export class HomeComponent {
  public shouldShowCreateProfile = false;
  public currentLoggedInUser: String = "";

  constructor (public auth: AuthService, @Inject(DOCUMENT) public document: Document, public userService: UserServiceService) {}

  // a double subscribe like this is probably not best practice, but  for now it works
  ngOnInit(): void {
    this.auth.user$.subscribe(result => {
      if (result) {
        console.log('the username: ', result.name);
        this.currentLoggedInUser = result.name as String;
        this.userService.getUserByUsername(this.currentLoggedInUser).subscribe(result => {
          console.log('searching current logged in user result: ', result);
          if (!result || !result.username) {
            this.shouldShowCreateProfile = true;
          }
        });
      }
    });
  }

  login(isSignUp: boolean) {
    this.auth.loginWithRedirect({
      appState: { 
        target: this.document.location.pathname
      },
      authorizationParams: {
        // Default selected login / signup
        screen_hint: isSignUp ? 'signup' : 'signin'
      }    
    });
    //this.showCreateProfile = isSignUp;
  }

  logout() {
    this.auth.logout({ 
      logoutParams: {
        returnTo: this.document.location.origin 
      }
    });
  }

  switchToHomeView(): void {
    this.shouldShowCreateProfile = false;
  }

  showCreateProfilePage(): void {
    this.shouldShowCreateProfile = true;
  }
}
