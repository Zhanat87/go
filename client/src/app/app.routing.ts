import { Routes, RouterModule } from '@angular/router';
import { ModuleWithProviders } from '@angular/core';
import {Login} from "./pages/login/login.component";

export const routes: Routes = [
  { path: 'login', component: Login },
  { path: '', redirectTo: 'index', pathMatch: 'full' },
    // @todo 404, 403, 500
  // { path: '**', redirectTo: 'pages/home' } // 404
];

export const routing: ModuleWithProviders = RouterModule.forRoot(routes); // , { useHash: true }
