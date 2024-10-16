import { DOCUMENT } from '@angular/common';
import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { environment } from '../../environment';
import { UserServiceService } from 'src/services/user-service.service';
import {MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import { map, Observable, startWith } from 'rxjs';
import { FormControl } from '@angular/forms';
import { PostService } from 'src/services/posts-service.service';
import { PostModel } from 'src/models/post-model';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css','../../styles.css'],
})
export class HomeComponent {
  public shouldShowCreateProfile = false;
  public shouldShowHomePage = false;
  public currentLoggedInUser: string = "";
  public isDevelopmentEnvironment = !environment.production;
  public searchTerm: string = "";
  public isLoggedIn: boolean = false;
  public friendList: string[] = [
    "Melissa Brown",
    "Avery Tribbett",
    "Cade Beckers",
    "Youssef Ibrahim"
  ];
  filteredOptions: Observable<string[]> = new Observable<string[]>();
  myControl = new FormControl('');
  public shouldShowProfilePage = false;
  public feed: PostModel[] = []

  constructor (public auth: AuthService, @Inject(DOCUMENT) public document: Document, public userService: UserServiceService, public postService: PostService) {}

  // a double subscribe like this is probably not best practice, but  for now it works
  ngOnInit(): void {
    this.auth.user$.subscribe(result => {
      // check if someone is signing up / logging in via auth0
      if (result) {
        this.currentLoggedInUser = result.name as string;
        this.isLoggedIn = true;

        // check for user in database
        this.userService.getUserByUsername(this.currentLoggedInUser).subscribe(result => {

          // if user does not exist in DB redirect to profile creation page else show home page
          if (!result || !result.username) {
            this.shouldShowCreateProfile = true;
          }
          else {
            this.shouldShowHomePage = true;
          }
        });

        this.postService.getInitialFeedByTime(5).subscribe(result => {
          if (result.length) {
            this.feed = result
          }
        })
      }

      // if no one logged in via auth0 default to home page view
      else {
        this.shouldShowHomePage = true;
      }
    });

    this.filteredOptions = this.myControl.valueChanges.pipe(
      startWith(''),
      map(value => this._filter(value || '')),
    );
  }

  private _filter(value: string): string[] {
    const filterValue = value.toLowerCase();

    // add something in here to async find what they have typed in on keydown, and return the list...?

    return this.friendList.filter(option => option.toLowerCase().includes(filterValue));
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
  }

  logout() {
    this.auth.logout({ 
      logoutParams: {
        returnTo: this.document.location.origin 
      }
    });
  }

  switchToHomeViewAndLogout(): void {
    this.shouldShowHomePage = true;
    this.shouldShowCreateProfile = false;
    this.shouldShowProfilePage = false;
    this.logout();
  }

  showCreateProfilePage(): void {
    this.shouldShowCreateProfile = true;
  }

  findFriend(searchString: any): void {
    // to add more functionality
  }

  showProfilePage(): void {
    this.shouldShowHomePage = false;
    this.shouldShowCreateProfile = false;
    this.shouldShowProfilePage = true;
  }

  switchToHomeView(): void {
    this.shouldShowHomePage = true;
    this.shouldShowCreateProfile = false;
    this.shouldShowProfilePage = false;
  }
}
