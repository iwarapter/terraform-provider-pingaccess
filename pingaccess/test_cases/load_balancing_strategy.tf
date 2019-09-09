resource "pingaccess_load_balancing_strategy" "demo_strategy" {
   name                         = "Round robin policy"
   class_name                    = "com.pingidentity.pa.ha.lb.roundrobin.CookieBasedRoundRobinPlugin"
   configuration     = << EOF
   {
        "stickySessionEnabled": true,
        "cookieName": "PA_S",
   }
   EOF
}