export class Post {
    public Id:number;
    public Title:string;
    public Description:string;
    public CategoryId:number;
    public CategoryName:string;

    constructor(id:number, title: string, description: string, categoryId:number, categoryName:string){
        this.Id = id;
        this.Title = title;
        this.Description = description;
        this.CategoryId = categoryId;
        this.CategoryName = categoryName;
    }
}
