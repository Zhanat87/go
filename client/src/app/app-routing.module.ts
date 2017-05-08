import { Routes, RouterModule } from '@angular/router';
import {NgModule} from '@angular/core';

import {AuthGuard} from "./common/services/auth.guard";

import { MainLayoutComponent } from "./common/components/layouts/main/main.component";
import { BlankLayoutComponent } from "./common/components/layouts/blank/blank.component";

// single pages
import {HomeComponent} from "./pages/home/home.component";
import {LoginComponent} from "./pages/login/login.component";
import { PageNotFoundComponent } from "./pages/404/page-not-found.component";
import {ChatComponent} from "./pages/chat/chat.component";
import {RegisterComponent} from "./pages/register/register.component";

export const routes: Routes = [

  // full layout
  {
    path: '',
    component: MainLayoutComponent,
    canActivateChild: [AuthGuard],
    children: [
      { path: '', redirectTo: 'index', pathMatch: 'full' },
      { path: 'index', component: HomeComponent },
      { path: 'chat', component: ChatComponent },
      { path: 'users', loadChildren: 'app/modules/user/user.module#UserModule' },
      { path: 'albums', loadChildren: 'app/modules/album/album.module#AlbumModule' },
      { path: 'artists', loadChildren: 'app/modules/artist/artist.module#ArtistModule' },
    ],
  },
  // simple layout
  {
    path: '',
    component: BlankLayoutComponent,
    children: [
      { path: 'login', component: LoginComponent },
      { path: 'register', component: RegisterComponent },
      {
        // must be at end of all routes
        path: '**',
        component: PageNotFoundComponent,
        data: {
          title: 'Page 404'
        }
      },
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  providers: [
    AuthGuard,
  ],
})
export class AppRoutingModule {}
