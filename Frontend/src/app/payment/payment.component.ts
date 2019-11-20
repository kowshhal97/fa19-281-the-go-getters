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
orderID:any

endpoint="http://54.90.233.215:80/payments"
  constructor(private http : HttpClient, private router: Router) { }

  ngOnInit() {
    this.orderID=sessionStorage.getItem('orderId')
    console.log(this.orderID)
  }
  pay(){

    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    
     this.http
        .post(this.endpoint,{
          "totalPrice":this.totalPrice,
          "cardDetails":this.cardnumber,
          "expDate":this.expDate,
          "contactPhone":this.phone,
          "securityCode":this.cvv
        },{headers: header})
        .subscribe((res) => {
            //do something with the response here

            this.placeOrderFinal()
            this.router.navigate(['./payment']);
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
        this.totalPrice=res['itemamount']+this.totalPrice
});
}

placeOrderFinal(){

  let header = new HttpHeaders();
  header.append('Content-Type', 'application/json');
  
   this.http
      .put("http://34.217.1.118:3000/order"+'/'+sessionStorage.getItem('orderId'),{headers: header}).subscribe((res) => {
          //do something with the response here
          this.router.navigate(['./payment']);
          
          console.log(res);

});
}

}
