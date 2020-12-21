import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {RouterModule, Routes} from '@angular/router';
import {UserComponent} from './pages/user/user.component';

const routes: Routes = [
  {
    path: '',
    children: [
      {
        path: '',
        component: UserComponent
      }
    ]
  }
];

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forChild(routes)
  ]
})
export class UserManagementModule { }
