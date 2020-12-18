import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LandingComponent} from './domain/lp/landing/landing.component';

const routes: Routes = [
  {
    path: 'auth',
    loadChildren: () => import('./domain/auth/auth.module').then(m => m.AuthModule)
  },
  {
    path: '',
    component: LandingComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
