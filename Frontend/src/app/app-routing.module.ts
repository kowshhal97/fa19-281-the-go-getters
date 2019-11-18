import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { LandingComponent } from './landing/landing.component';
import { RegistrationComponent } from './registration/registration.component';




const appRoutes: Routes = [
];

@NgModule({
  imports: [RouterModule.forRoot(appRoutes, { useHash: false })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
