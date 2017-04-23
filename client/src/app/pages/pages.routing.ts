import { Routes, RouterModule }  from '@angular/router';
import { Pages } from './pages.component';
import { ModuleWithProviders } from '@angular/core';
import {AuthGuard} from "../common/services/auth.guard";
// noinspection TypeScriptValidateTypes

// export function loadChildren(path) { return System.import(path); };

export const routes: Routes = [
  // {
  //   path: 'register',
  //   loadChildren: 'app/pages/register/register.module#RegisterModule'
  // },
  {
    path: '',
    component: Pages,
    // canActivateChild: [AuthGuard],
    children: [
      { path: '', redirectTo: 'index', pathMatch: 'full' },
      { path: 'index', loadChildren: 'app/pages/home/home.module#HomeModule' },
      { path: 'users', loadChildren: 'app/modules/user/user.module#UserModule' },
      { path: 'albums', loadChildren: 'app/modules/album/album.module#AlbumModule' },
      { path: 'artists', loadChildren: 'app/modules/artist/artist.module#ArtistModule' },

      // { path: 'dashboard', loadChildren: 'app/pages/dashboard/dashboard.module#DashboardModule' },
      // { path: 'editors', loadChildren: 'app/pages/editors/editors.module#EditorsModule' },
      // { path: 'components', loadChildren: 'app/pages/components/components.module#ComponentsModule' },
      // { path: 'charts', loadChildren: 'app/pages/charts/charts.module#ChartsModule' },
      // { path: 'ui', loadChildren: 'app/pages/ui/ui.module#UiModule' },
      // { path: 'forms', loadChildren: 'app/pages/forms/forms.module#FormsModule' },
      // { path: 'tables', loadChildren: 'app/pages/tables/tables.module#TablesModule' },
      // { path: 'maps', loadChildren: 'app/pages/maps/maps.module#MapsModule' }
    ]
  }
];

export const routing: ModuleWithProviders = RouterModule.forChild(routes);
