import { DOCUMENT } from '@angular/common';
import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css','../../styles.css'],
})
export class HomeComponent {
  constructor (public auth: AuthService, @Inject(DOCUMENT) public document: Document) {}

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
  }

  logout() {
    this.auth.logout({ 
      logoutParams: {
        returnTo: this.document.location.origin 
      }
    });
  }
}
