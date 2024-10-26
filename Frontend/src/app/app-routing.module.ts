import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { UserProfileComponent } from './user-profile/user-profile.component';
import { UserFeedComponent } from './user-feed/user-feed.component';
import { CreateProfileComponent } from './create-profile/create-profile.component';
import { DeveloperPageComponent } from './developer-page/developer-page.component';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  {
    path: 'profile/:profileUser/:isDeveloperMode',
    component: UserProfileComponent,
  },
  { path: 'feed', component: UserFeedComponent },
  {
    path: 'create-profile/:currentUser/:isDeveloperMode',
    component: CreateProfileComponent,
  },
  {
    path: 'developer/:isDevelopmentEnvironment',
    component: DeveloperPageComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
