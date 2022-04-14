#include<stdio.h>
int main()
{
   /* int i,j,n;
    scanf("%d",&n);
    for(i=1;i<=n;i++)
    {
        for(j=1;j<=i;j++)
        {
            printf("i ");
        }
        printf("\n");
    }*/
    /*int arr[100],value,n,i,j,temp;
    scanf("%d",&n);
    for(i=0;i<n;i++)
    {
        scanf("%d",&value);
        arr[i]=value;
    }
    //bubble sort
    for(i=0;i<n-1;i++)//12 2 3 7 1 i=0

    {
        for(j=i+1;j<n;j++) //j=1,2
        {
            if(arr[i]>arr[j]) //arr[0]=2, arr[2]=3
            {
               //swap
                temp=arr[j];//temp=2
                arr[j]=arr[i];//arr[1]=12
                arr[i]=temp;//arr[0]=2
            }
        }
        //2 12 3 7 1
    }
    for(i=0;i<n;i++)
    {
        printf("%d ",arr[i]);
    }*/

   char str[25];
   int i,count =1;
   printf("Enter the string: ");
   gets(str);

   for(i=0;i<=strlen(str);i++){
      /*if(str[i]>=65&&str[i]<=90)
         str[i]=str[i]+32;
    else if(str[i]>=97&&str[i]<=122)
    {
        str[i]=str[i]-32;
    }
   else{
      continue;*/
      if(str[i]==32)
        count++;
   }
    printf("Total Words:%d",count);

   //printf("updated String is: %",str);


 return 0;
}
