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

endpoint="https://i18253eej8.execute-api.us-east-1.amazonaws.com/prod/payments"
  constructor(private http : HttpClient, private router: Router) { }

  ngOnInit() {
    if(sessionStorage.getItem('userId')==null)
    {
      this.router.navigate(['./login'])
      window.alert("you need to login first!")
    }
    if(sessionStorage.getItem('orderId')==null)
    {
      this.router.navigate(['./order'])
      window.alert("you need to order first!")
    }
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
      .put("https://i18253eej8.execute-api.us-east-1.amazonaws.com/prod/order"+'/'+sessionStorage.getItem('orderId'),{headers: header}).subscribe((res) => {
          //do something with the response here
          this.router.navigate(['./order']);
          
          console.log(res);

});
}
gotoMenu(){
  this.router.navigate(['./menu'])
}

gotoReviews(){
  this.router.navigate(['./reviews'])
}

gotoHome(){
  this.router.navigate(['./home'])
}
logout(){
  sessionStorage.setItem('userId',null)
  this.router.navigate(['./login'])
}
}
