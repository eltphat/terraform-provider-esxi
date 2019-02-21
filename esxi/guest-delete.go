package esxi

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
<<<<<<< HEAD
=======
	"time"
>>>>>>> a09975692ab4114aef08427f9b410b63842981c3
)

func resourceGUESTDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	esxiSSHinfo := SshConnectionStruct{c.esxiHostName, c.esxiHostPort, c.esxiUserName, c.esxiPassword}
<<<<<<< HEAD
=======
	log.Println("[resourceGUESTDelete]")
>>>>>>> a09975692ab4114aef08427f9b410b63842981c3

	var remote_cmd, stdout string
	var err error

	vmid := d.Id()
	guest_shutdown_timeout := d.Get("guest_shutdown_timeout").(int)

	_, err = guestPowerOff(c, vmid, guest_shutdown_timeout)
	if err != nil {
		return err
	}

	// remove storage from vmx so it doesn't get deleted by the vim-cmd destroy
	err = cleanStorageFromVmx(c, vmid)
	if err != nil {
		log.Printf("[resourceGUESTDelete] Failed clean storage from vmid: %s (to be deleted)\n", vmid)
	}

<<<<<<< HEAD
=======
	time.Sleep(5 * time.Second)
>>>>>>> a09975692ab4114aef08427f9b410b63842981c3
	remote_cmd = fmt.Sprintf("vim-cmd vmsvc/destroy %s", vmid)
	stdout, err = runRemoteSshCommand(esxiSSHinfo, remote_cmd, "vmsvc/destroy")
	if err != nil {
		// todo more descriptive err message
		log.Printf("[resourceGUESTDelete] Failed destroy vmid: %s\n", stdout)
		return err
	}

	d.SetId("")
<<<<<<< HEAD
=======

>>>>>>> a09975692ab4114aef08427f9b410b63842981c3
	return nil
}
