# message_generator


Each directory in **message_generator** represents a message

### This package allows you to generate a message and have it *directly copied* to your clipboard.

For example:

> $ cd billingOrderItemCancelled/
> 
> $ go run main.go AccountID 12345 OrderID 5555

*The arguments are organized by <field_name> <field_value> and always are in pairs.* 

*The timestamp is accurate to whenever you execute the above command*

### This command will copy the message to your clipboard

> {"meta":{"timestamp":"2023-04-17T16:13:39.754443-06:00"},"account_id":12345,"order_id":5555}


### To generate aliases  

> cd message_generator
> go run build.go

This adds an alias to your `.zshrc` file for every directory in the message generator.

*The names of the aliases are the same names as the directories*

For Example:

> $ billingOrderItemCancelled AccountID 12345 OrderID 5555
