import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { LandingComponent } from './landing/landing.component';
import { RegistrationComponent } from './registration/registration.component';
import { OrderComponent } from './order/order.component';
import { MenuComponent } from './menu/menu.component';
import { ReviewsComponent } from './reviews/reviews.component';
import { PaymentComponent } from './payment/payment.component';




const appRoutes: Routes = [
  { path: '', component: LoginComponent, pathMatch: 'full'},
  { path: 'register', component: RegistrationComponent },
  { path: 'login', component: LoginComponent},
  { path: 'home', component: LandingComponent},
  { path: 'order', component:OrderComponent},
  { path: 'menu', component:MenuComponent},
  { path: 'reviews', component:ReviewsComponent},
  {path: 'payment', component:PaymentComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(appRoutes, { useHash: false })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
