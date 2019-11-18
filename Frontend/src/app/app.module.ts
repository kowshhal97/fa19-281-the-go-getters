import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule,Router,Routes } from '@angular/router';


import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';

import { HomeComponent } from './home/home.component';
import { RegistrationComponent } from './registration/registration.component';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { LandingComponent } from './landing/landing.component';

const myRoots: Routes = [
  { path: '', component: LoginComponent, pathMatch: 'full'},
  { path: 'register', component: RegistrationComponent },
  { path: 'login', component: LoginComponent},
  { path: 'home', component: LandingComponent}
];

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,

    HomeComponent,
    RegistrationComponent,
    LandingComponent
  ],
  imports: [
    BrowserModule,ReactiveFormsModule,FormsModule,
    RouterModule.forRoot(
      myRoots,
      { enableTracing: true } // <-- debugging purposes only
    )
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }