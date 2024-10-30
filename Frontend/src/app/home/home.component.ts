import { DOCUMENT } from '@angular/common';
import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { environment } from '../../environment';
import { UserServiceService } from 'src/services/user-service.service';
import { Observable } from 'rxjs';
import { FormControl, FormGroup } from '@angular/forms';
import { DisplayNameUserModel, UserModel } from 'src/models/user-model';
import { Router } from '@angular/router';
import { PostService } from 'src/services/posts-service.service';
import { PostModel } from 'src/models/post-model';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css', '../../styles.css'],
})
export class HomeComponent {
  public currentLoggedInUsername = '';
  public currentSearchOptions: DisplayNameUserModel[] = [];
  public currentUser = {
    id: 0,
    firstName: 'Please choose a name...',
    lastName: '',
    username: '',
  } as UserModel;
  public isDevelopmentEnvironment = !environment.production;
  public searchTerm = '';
  public isLoggedIn = false;
  public friendSearchList: DisplayNameUserModel[] = [];
  public friendList: DisplayNameUserModel[] = [];
  filteredOptions: Observable<DisplayNameUserModel[]> = new Observable<
    DisplayNameUserModel[]
  >();
  myControl = new FormControl('');
  public shouldShowProfilePage = false;
  public feed: PostModel[] = [];
  public numOfPosts = 0;
  public isLoading = false;
  public endOfFeed = false;
  public createPostForm: FormGroup = new FormGroup({});

  constructor(
    public auth: AuthService,
    @Inject(DOCUMENT) public document: Document,
    public userService: UserServiceService,
    public router: Router,
    public postService: PostService,
  ) {
    this.document.addEventListener('scroll', this.onScroll.bind(this));
  }

  // a double subscribe like this is probably not best practice, but  for now it works
  ngOnInit(): void {
    this.userService.userAuth.user$.subscribe((result) => {
      // check if someone is signing up / logging in via auth0
      if (result) {
        this.currentLoggedInUsername = result.name as string;
        this.userService.setValue(result.name as string);
        this.isLoggedIn = true;

        // check for user in database
        this.userService
          .getUserByUsername(this.currentLoggedInUsername)
          .subscribe((result) => {
            // if user does not exist in DB redirect to profile creation page else show home page
            if (!result || !result.username) {
              this.router.navigate(['/create-profile', result.username, false]);
            } else {
              this.router.navigate(['/home']);
              this.currentUser = result;
            }
          });
        this.postService.getInitialFeedByTime(20).subscribe((result) => {
          if (result.length) {
            this.feed = result;
          }
        });
      }
    });

    this.myControl.valueChanges.subscribe((searchValue) => {
      const newVal = searchValue as string;
      if (newVal) {
        this.userService.searchUser(newVal).subscribe((result) => {
          const userList: DisplayNameUserModel[] = [];
          if (result) {
            for (const user of result) {
              userList.push({
                firstName: user.firstName,
                lastName: user.lastName,
              });
            }
            this.friendList = userList;
          }
        });
      }
    });

    this.initPostForm();
  }

  public initPostForm(): void {
    this.createPostForm = new FormGroup({
      createPostInput: new FormControl(''),
    });
  }

  login(isSignUp: boolean) {
    this.userService.userAuth.loginWithRedirect({
      appState: {
        target: this.document.location.pathname,
      },
      authorizationParams: {
        // Default selected login / signup
        screen_hint: isSignUp ? 'signup' : 'signin',
      },
    });
  }

  logout() {
    this.userService.userAuth.logout({
      logoutParams: {
        returnTo: this.document.location.origin,
      },
    });
  }

  findFriend(searchFirstName: string, searchLastName: string): void {
    this.userService
      .searchUserByFirstAndLastName(searchFirstName, searchLastName)
      .subscribe((result) => {
        this.router.navigate(['/profile', result.username, false]);
      });
  }

  // Method to handle scroll event
  onScroll(): void {
    // Calculate the distance from the bottom of the page
    const scrollPosition = window.scrollY;
    const threshold = document.body.scrollHeight - 1000; // Adjust threshold as needed

    // Check if the user has scrolled near the bottom and if not currently loading
    if (scrollPosition >= threshold && !this.isLoading && !this.endOfFeed) {
      this.loadMorePosts();
    }
  }

  // Method to load more posts
  async loadMorePosts(): Promise<void> {
    this.isLoading = true;
    this.numOfPosts += 20;

    // Fetch 20 more posts from backend
    this.postService.getFeedByTime(this.numOfPosts).subscribe(
      async (result: PostModel[]) => {
        // Check for success
        if (result && result.length) {
          // Pause for 1 seconds to simulate loading (api takes like 1 milisecond)
          await new Promise((r) => setTimeout(r, 1000));
          // Append new posts to existing feed
          this.feed = [...this.feed, ...result];
        } else {
          // Once we get an empty array we know we are at the end of the feed
          this.endOfFeed = true;
        }
        // Loading is complete
        this.isLoading = false;
      },
      (error) => {
        // Will just be the end of the feed
        this.isLoading = false;
        console.error('Error loading posts:', error);
      },
    );
  }

  createNewPost(): void {
    if (!this.currentUser.id) {
      return;
    }
    this.postService
      .addPost(
        this.currentUser.id,
        this.createPostForm.get('createPostInput')?.value,
      )
      .subscribe((result) => {
        if (result) {
          // If successful post reinitialize feed
          this.isLoading = true;
          this.postService.getInitialFeedByTime(20).subscribe((result) => {
            // Reset component
            if (result.length) {
              this.feed = result;
              this.numOfPosts = 0;
              this.createPostForm = new FormGroup({
                createPostInput: new FormControl(''),
              });
            }
          });
          this.isLoading = false;
        } else {
          console.error('Something went wrong making the post');
        }
      });
  }
}
