import { Component, OnInit } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-payment',
  templateUrl: './payment.component.html',
  styleUrls: ['./payment.component.css']
})
export class PaymentComponent implements OnInit {
totalPrice:any
cardnumber:any
//expDate:any
phone:any
cvv:any
expMonth:any
expyear:any
expDate=this.expMonth+'/'+this.expyear;

  constructor(private http : HttpClient, private router: Router) { }

  ngOnInit() {
  }
  signup(){

    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    
     this.http
        .post("endpoint not available yet",
        {"totalPrice":this.totalPrice,"cardDetails":this.cardnumber,"expDate":this.expDate,"contactPhone":this.phone,"securityCode":this.cvv},
            {headers: header})
        .subscribe((res) => {
            //do something with the response here

            this.router.navigate(['./home']);

            console.log(res);

});
}
getItemById(){

  let header = new HttpHeaders();
  header.append('Content-Type', 'application/json');
  
   this.http
      .post("endpoint not available yet",
      {"totalPrice":this.totalPrice,"cardDetails":this.cardnumber,"expDate":this.expDate,"contactPhone":this.phone,"securityCode":this.cvv},
          {headers: header})
      .subscribe((res) => {
          //do something with the response here

          this.router.navigate(['./home']);
          
          console.log(res);

});
}
}
