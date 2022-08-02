import { Category } from './../model/category';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class CategoryService {
  endpoint: string = "http://localhost:8081"
  
  constructor(private httpClient: HttpClient) {}

  getAllCategories(): Observable<any>{
    return this.httpClient.get<Category[]>(this.endpoint+ "/category")
  }
  getOneById(id:number):Observable<Category[]>{
    return this.httpClient.get<Category[]>(this.endpoint+ "" + id)
  }

  deleteCategory(id:number){
    return this.httpClient.delete(this.endpoint + '' + id, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
      },
    });
  }

  createCategory(category: Category){
    return this.httpClient.post(this.endpoint + '/category', category, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
      },
    });
  }
}
