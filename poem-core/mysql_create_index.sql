ALTER TABLE `poem`.`t_book`
ADD INDEX `IDX_t_book_Brief` (`Brief`(50) ASC) ;

ALTER TABLE `poem`.`t_book_item`
ADD INDEX `IDX_t_book_item_Content` (`Content`(50) ASC) ;

ALTER TABLE `poem`.`t_poem`
ADD INDEX `IDX_t_poem_Content` (`Content`(50) ASC) ;
ALTER TABLE `poem`.`t_poem`
ADD INDEX `IDX_t_poem_Shang` (`Shang`(50) ASC) ;

ALTER TABLE `poem`.`t_poer`
ADD INDEX `IDX_t_poer_Brief` (`Brief`(50) ASC) ;