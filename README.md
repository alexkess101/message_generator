# message_generator


Each directory in **message_generator** represents a message

### This package allows you to generate a message and have it *directly copied* to your clipboard.

For example:

> $ cd billingOrderItemCancelled/
> 
> $ go run main.go AccountID 12345 OrderID 5555


**This command will copy the message to your clipboard*

> {"meta":{"timestamp":"2023-04-17T16:13:39.754443-06:00"},"account_id":12345,"order_id":5555}

The arguments are organized by <field_name> <field_value> and always are in pairs.

The timestamp is accurate to whenever you execute the above command

### Installation
Run `make install` in the project directory

