select BillingCountry, sum(total) as Res from Invoice group by BillingCountry order by Res ASC;
