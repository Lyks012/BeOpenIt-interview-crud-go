import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Post } from '../model/post';

@Injectable({
  providedIn: 'root',
})
export class PostsService {
  endpoint: string = 'http://localhost:8081';

  constructor(private httpClient: HttpClient) {}

  getAllPosts(): Observable<any> {
    return this.httpClient.get<any>(this.endpoint + '/posts');
  }
  getByCategoryId(id: number): Observable<Post[]> {
    return this.httpClient.get<Post[]>(
      this.endpoint + '/category/' + id + '/posts'
    );
  }

  getOneById(id: number): Observable<Post[]> {
    return this.httpClient.get<Post[]>(this.endpoint + '/post:' + id);
  }

  deletePost(id: number) {
    return this.httpClient.delete(this.endpoint + '/post/' + id, {
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
      },
    });
  }

  createPost(newPost: Post) {
    return this.httpClient.post(
      this.endpoint + '/createPost',
      {
        Id: 0,
        Title: newPost.Title,
        Description: newPost.Description,
        Category_id: newPost.CategoryId,
      },
      {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8',
        },
      }
    );
  }
}
