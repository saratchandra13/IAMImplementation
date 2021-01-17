# IAMImplementation
Sample IAM implementation with few functionalities


Added sample IAM implementation by supporting following use-cases.

1. check if user is eligible to access a resource with a given permission.
2. create a role which would have multiple resources and permissions attatched to it.
3. CRUD for users, roles and resources.


Extensions.

Here there are 3 variables.

1. resources 2. users 3. permissions

If there are multiple resources we should create resource groups for easier access and storage.
If there are multiple users we should create user group and add user to it.
If there are multiple permissions we should create roles (like admin, developer etc..) this is implemented in this.

