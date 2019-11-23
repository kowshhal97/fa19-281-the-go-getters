import { Component, OnInit, Input } from '@angular/core';
import { Router } from '@angular/router';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { bindCallback } from 'rxjs';

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
  endpoint="https://i18253eej8.execute-api.us-east-1.amazonaws.com/prod/order"
  menu:any
  constructor(private http : HttpClient, private router: Router) {
    
  }
  totalAmount:any
  orderstatus:any
  orderresponseObject:any
  itemPriceee: any; 

  orderDetails=[]
  ngOnInit() {
    if(sessionStorage.getItem('userId')==null)
    {
      this.router.navigate(['./login'])
      window.alert("you need to login first!")
    }

  }
  setOrderObject(){
    this.orderresponseObject={
      'orderId':this.orderId,
      'itemName':this.itemName,
      'totalAmount':this.totalAmount,
      'orderStatus':this.orderstatus
    }
  }
  getMenuItem(){
    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    const p = this.http.get("https://i18253eej8.execute-api.us-east-1.amazonaws.com/prod/menu"+'/'+this.itemId,
            {headers: header})
        .subscribe((res) => {
          this.itemPrice=res['Price']
          this.itemName=res['ItemName']
            //do something with the response here
            console.log(res)
            this.router.navigate(['./order']);
        });

    
  }
  
  placeOrder() {
    
    let header = new HttpHeaders();
    header.append('Content-Type', 'application/json');
    const p = this.http.get("https://i18253eej8.execute-api.us-east-1.amazonaws.com/prod/menu"+'/'+this.itemId,
            {headers: header})
        .subscribe((res) => {
          this.itemPrice=res['Price']
          this.itemName=res['ItemName']
            //do something with the response here
            console.log(res)
            this.router.navigate(['./order']);
            header.append('Content-Type', 'application/json');
            this.http.post('https://0ghg7tvccf.execute-api.us-west-2.amazonaws.com/prod/order',JSON.stringify({"userId" : this.userid,
               "itemName" : this.itemName,
               "itemQuantity" : parseInt(this.itemQuantity),
               "itemPrice": parseFloat(this.itemPrice),
               "itemId" : this.itemId}),{headers: header}).subscribe((res) => {
                 this.userid =sessionStorage.getItem("userId");
                   //do something with the response here
                   this.orderId=res['orderId']
                   this.totalAmount=res['totalAmount']
                   this.orderstatus=res['orderStatus']
                   sessionStorage.setItem('totalAmount',this.totalAmount)
                   this.setOrderObject()
                   this.orderDetails=[this.orderresponseObject]
                   sessionStorage.setItem('orderId',this.orderId)
                   this.router.navigate(['./order']);
                   //window.alert("Order Placed, Id="+this.orderId)
                    console.log(res);
               });            
        });
    //let header = new HttpHeaders();
  
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
          this.orderId=res['orderId']
            this.totalAmount=res['totalAmount']
            this.orderstatus=res['orderStatus']
            this.setOrderObject()
            this.orderDetails=[this.orderresponseObject]
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

