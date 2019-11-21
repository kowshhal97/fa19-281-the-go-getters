import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import { HttpHeaders, HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-order',
  templateUrl: './order.component.html',
  styleUrls: ['./order.component.css']
})
export class OrderComponent implements OnInit {


userid: any;
itemName:any
itemQuantity: any
itemPrice:any
itemId:any
getOrderId:any
deleteOrderId:any
orderId:any=""
getorder:any
  endpoint="http://52.27.19.100:3000/order"
  menu:any
  constructor(private http : HttpClient, private router: Router) {
    
  }
  ngOnInit() {
    if(sessionStorage.getItem('userId')==null)
    {
      this.router.navigate(['./login'])
      window.alert("you need to login first!")
    }

  }
  getMenuItem(){
    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
     this.http.get("http://52.12.73.70:8001/menu"+'/'+this.itemId,
            {headers: header})
        .subscribe((res) => {
          this.itemPrice=res['Price']
          console.log(this.itemPrice)
          this.itemName=res['itemName']
            //do something with the response here
            this.router.navigate(['./order']);
            console.log(res);
        });
  }
  placeOrder() {
    this.getMenuItem()
    console.log(this.itemPrice)
    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
     this.http.post(this.endpoint,{"userId" : this.userid,
        "itemName" : this.itemName,
        "itemQuantity" : parseInt(this.itemQuantity),
        "itemPrice": parseFloat(this.itemPrice),
        "itemId" : this.itemId},{headers: header}).subscribe((res) => {
          this.userid =sessionStorage.getItem("userId");
            //do something with the response here
            this.orderId=res['orderId']
            sessionStorage.setItem('orderId',this.orderId)
            this.router.navigate(['./order']);
            window.alert("Order Placed, Id="+this.orderId)
            console.log(res);
        }); 
  }
  goToPayments(){
    this.router.navigate(['./payment'])
  }
  cancelOrder(){
      let header = new HttpHeaders();
      header.append('Content-Type', 'application/json');
      
       this.http
          .delete(this.endpoint+'/'+this.orderId,
              {headers: header})
          .subscribe((res) => {
              //do something with the response here
              this.router.navigate(['./order']);
              console.log(res);
            });  
  }

  getOrder(){
    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
     this.http.get(this.endpoint+'/'+this.getOrderId,
            {headers: header})
        .subscribe((res) => {
          this.getorder=res
            //do something with the response here
            this.router.navigate(['./order']);
            console.log(res);
        });
  }
  deleteOrder(){
    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    
     this.http
        .delete(this.endpoint+'/'+this.deleteOrderId,
            {headers: header})
        .subscribe((res) => {
          window.alert("order deleted")
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

